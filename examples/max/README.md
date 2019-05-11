# Go Template Example

## 如何啟動

```
make run
```

訪問: http://localhost:3000

## 知識點

#### 1. func New

```
func New(name string) *Template
```

`template.New` 生成一個新的 Template 可以接著用 template 的 methods。

*範例:*

```go
template.New("index.html").<template methods>
```

#### 2. func (*Template) Funcs

```
func (t *Template) Funcs(funcMap FuncMap) *Template
```

為 templeate 添加一些函數。 接收 `FuncMap` 類型；必須要在 parse template 前使用。

搭配 `template.FuncMap` 使用。

*範例:*

```go
newTemplate := template.New("index.html")

newTemplate.Funcs(template.FuncMap{
        "templateMethod1": isOdd,
        "templateMethod2": isEven,
        "templateMethod2": isError,
        // ...
    })
```

```html
<!-- 在 html 調用 template function -->
{{if isError data }}
    <span>Something went wrong</span>
{{end}}
```

#### 3. func (*Template) ParseFiles

```
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

解析 template , 可以放入多個 template。

*範例:*

```go
templates := []string{
    "web/template/layout/base.html",
    "web/template/layout/footer.html",
    "web/template/layout/header.html",
    "web/template/index.html",
}

template.New("index.html").ParseFiles(templates...)
```

#### 4. func (*Template) ExecuteTemplate

將數據關聯到給予的模板檔案名并寫到指定的 `io.Writer` 裡面

*範例:*

```go
t, _ := template.New("index.html").
        Funcs(template.FuncMap{"isOdd": isOdd}).
        ParseFiles(templates...)

// 將數據給予 base.html
t.ExecuteTemplate(w, "base", items)
```

```html
<!-- 使用數據 -->
{{range $i, $v := .}}
<tr {{if isOdd $i}} class="odd" {{end}}>
    <td>{{ $v.Name }}</td>
    <td>{{ $v.Work }}</td>
    <td>{{ $v.Amateur }}</td>
</tr>
{{end}}
```

```html
<!-- 或者往下傳遞 -->
{{template "content" .}}
```


