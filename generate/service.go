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

package generate

import (
	"github.com/zeromicro/goctl-android/template"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/tools/goctl/util"
)

func generateService(dir string, data IService) error {
	filename := filepath.Join(dir, "Service.java")
	base := filepath.Dir(filename)
	err := os.MkdirAll(base, os.ModePerm)
	if err != nil {
		return err
	}

	return util.With("service").Parse(template.Service).SaveTo(data, filename, true)
}
