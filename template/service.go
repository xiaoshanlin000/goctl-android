// MIT License
//
// Copyright (c) 2020 goctl-android
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

package template

var Service = `package {{.ParentPackage}}.service;

{{.Import}}
import com.google.gson.Gson;
import okhttp3.MediaType;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class Service {
    private static final String MEDIA_TYPE_JSON = "application/json; charset=utf-8";
    private static final String BASE_RUL = "{{.Hostname}}";// TODO replace to your host and delete this comment
    private static IService service;
    private final Gson gson = new Gson();

    private Service() {
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(BASE_RUL)
                .addConverterFactory(GsonConverterFactory.create())
                .build();
        service = retrofit.create(IService.class);
    }

    private static class ServiceImpl{
        private static final Service instance = new Service();
    }

    public static Service getInstance() {
        return ServiceImpl.instance;
    }

    private RequestBody buildJSONBody(Object obj) {
        String s = gson.toJson(obj);
        return RequestBody.create(MediaType.parse(MEDIA_TYPE_JSON), s);
    }
	{{range $index,$item := .Routes}}{{$item.Doc}}
    public void {{$item.MethodName}}({{if $item.HasRequest}}{{$item.RequestBeanName}} in, {{end}}Callback{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} callback) {
        Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} call = service.{{$item.MethodName}}({{if $item.HavePath}}{{$item.PathId}}{{end}}{{if $item.HaveQuery}}{{$item.QueryId}}{{end}}{{if $item.ShowRequestBody}}buildJSONBody(in){{end}});
        call.enqueue(callback);
    }
	{{end}}
}
`
