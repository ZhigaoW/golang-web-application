

### Form


```html
    <form action="/process" method="post" enctype="application/x-www-form-urlencoded">
        <input type="text" name="first_name" />
        <input type="text" name="last_name" />
        <input type="submit" />
    </form>
```

- 传送的数据类型通过enctype决定
  - application/x-www-form-urlencoded : 数据会背编码为url查询字符串, 适合小文件
  - "multipart/form-data": 适合大文件
- 表单的method
  - get : 所有的数据都是通过url提取
  - post 



Request中的函数允许我们从url或/和Body中提取数据，通过这些字段

- Form
- PostForm
- MultipartForm

Form里面的数据是key-value对。 

通常的做法是
- 先调用Form或ParseMultiForm来解析Request
- 然后相应的访问Form、PostForm或MultipartForm字段
