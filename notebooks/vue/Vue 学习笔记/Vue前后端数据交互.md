# Vue前后端数据交互

## 一.前后端数据交互

下面是一个简单的前后端数据交互流程：

1. 在前端（浏览器）中发起HTTP请求：使用前端框架（如Vue.js）或者原生JavaScript，您可以通过发送HTTP请求（如GET、POST、PUT、DELETE）向后端服务器请求数据。
2. 后端服务器处理请求：后端服务器接收到前端发送的HTTP请求，并根据请求的类型和参数进行处理。根据具体业务需求，后端可能会从数据库中检索数据、执行某些计算或操作，并准备好要返回给前端的数据。
3. 后端返回响应数据：后端服务器将处理后的数据封装成HTTP响应，并发送回前端。通常，响应数据使用某种数据格式（如JSON）进行编码。
4. 前端接收响应数据：前端（浏览器）收到后端返回的HTTP响应后，解析响应数据。使用前端框架或原生JavaScript，您可以提取所需的数据。
5. 前端数据渲染：一旦前端获取到数据，您可以将其应用于您的用户界面，进行数据渲染和展示。这可以通过使用模板引擎、Vue组件或JavaScript操作DOM等方式来实现。

## 二.前端调用接口的方式

因为Vue不操作DOM，所以用后两种

- ajax :
- jQuery的ajax
- **fetch**
- **axios**

## 三.使用Flask 搭建后端服务

首先安装Flask 框架:

```python
pip install Flask
```

然后新建`data.json`文件:

```json
[
  {
    "id": "1",
    "name": "Alice",
    "age": 25
  },
  {
    "id": "2",
    "name": "Bob",
    "age": 30
  },
  {
    "id": "3",
    "name": "Charlie",
    "age": 35
  }
]
```

为了解决前端提交请求时出现的跨域问题，在Flask应用程序中添加CORS（跨域资源共享）中间件。这样可以在服务器端处理跨域请求的头信息，而无需在每个路由处理函数中手动添加跨域头。

这里使用Flask-CORS扩展来简化这个过程。首先，确保已经安装了Flask-CORS扩展包。可以使用以下命令来安装：

```bash
pip install flask-cors
```

以下是使用Flask框架实现对JSON文件进行增删改查操作的后端代码:

```python
from flask import Flask, request, jsonify
from flask_cors import CORS
import json

app = Flask(__name__)
CORS(app)

# 读取JSON文件
def read_json_file(file_path):
    with open(file_path, 'r') as file:
        data = json.load(file)
    return data

# 编辑JSON文件
def write_json_file(file_path, data):
    with open(file_path, 'w') as file:
        json.dump(data, file, indent=4)

# 获取所有数据
@app.route('/data', methods=['GET'])
def get_all_data():
    json_data = read_json_file('data.json')
    response = jsonify(json_data)
    return response

# 获取单个数据
@app.route('/data/<id>', methods=['GET'])
def get_data(id):
    json_data = read_json_file('data.json')
    for item in json_data:
        if item['id'] == id:
            response = jsonify(item)
            return response
    response = jsonify({'message': 'Data not found'})
    return response

# 添加数据
@app.route('/data', methods=['POST'])
def add_data():
    json_data = read_json_file('data.json')
    new_data = {
        'id': request.json['id'],
        'name': request.json['name'],
        'age': request.json['age']
    }
    json_data.append(new_data)
    write_json_file('data.json', json_data)
    response = jsonify({'message': 'Data added successfully'})
    return response

# 更新数据
@app.route('/data/<id>', methods=['PUT'])
def update_data(id):
    json_data = read_json_file('data.json')
    for item in json_data:
        if item['id'] == id:
            item['name'] = request.json.get('name', item['name'])
            item['age'] = request.json.get('age', item['age'])
            write_json_file('data.json', json_data)
            response = jsonify({'message': 'Data updated successfully'})
            return response
    response = jsonify({'message': 'Data not found'})
    return response

# 删除数据
@app.route('/data/<id>', methods=['DELETE'])
def delete_data(id):
    json_data = read_json_file('data.json')
    for item in json_data:
        if item['id'] == id:
            json_data.remove(item)
            write_json_file('data.json', json_data)
            response = jsonify({'message': 'Data deleted successfully'})
            return response
    response = jsonify({'message': 'Data not found'})
    return response

if __name__ == '__main__':
    app.run()
```



## 四、基于jQuery的ajax前后端交互模式

基于jQuery的Ajax（Asynchronous JavaScript and XML）是一种前后端交互模式，它利用JavaScript和XML（现在也可以使用JSON）来在Web应用程序中进行异步通信。它允许通过在不刷新整个页面的情况下从服务器获取数据和更新页面的内容。

使用基于jQuery的Ajax，可以通过以下步骤实现前后端交互：

1. 引入jQuery库：首先，在HTML页面中引入jQuery库文件，这样就可以使用jQuery的函数和方法。

```js
  <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
```

1. 编写Ajax请求：使用jQuery的`$.ajax()`方法来发送异步请求到服务器。该方法接受一个包含各种选项的JavaScript对象作为参数，用于配置请求的细节，例如URL、请求类型（GET、POST等）、数据等。

```js
$.ajax({
  url: 'example.com/api/data',
  type: 'GET',
  data: { param1: 'value1', param2: 'value2' },
  success: function(response) {
    // 处理成功响应的回调函数
    console.log(response);
  },
  error: function(xhr, status, error) {
    // 处理错误响应的回调函数
    console.log('Error:', error);
  }
});
```

在上面的代码中，我们指定了请求的URL、类型为GET，还传递了一些数据。成功回调函数在服务器响应成功时触发，错误回调函数在出现错误时触发。

1. 处理服务器响应：根据服务器返回的数据，可以在成功回调函数中执行一些操作。例如，更新页面内容、显示数据等。在上面的示例中，我们使用`console.log()`将响应数据打印到浏览器的控制台。

通过这种方式，前端可以通过Ajax向后端发送请求并处理响应，从而实现动态更新页面内容，而无需刷新整个页面。这种交互模式在创建交互性强、用户体验良好的Web应用程序时非常有用。

### 实现增删改查页面

```html
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <link href="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
  <script src="vue.js"></script>
</head>

<body>
<div class="container-fluid">
  <div class="row">
    <div class="col-md-6 col-md-offset-3" style="margin-top: 20px">
      <div id="app">
        <button @click="handleLoad" v-if="!dataLoaded">加载数据</button>
        <button @click="handleAdd">添加数据</button>
        <button @click="handleUpdate">更新数据</button>
        <button @click="handleDelete">删除数据</button>
        <hr>
        <div v-for="item in items" :key="item.id">
          <div>ID: {{ item.id }}</div>
          <div>名称: {{ item.name }}</div>
          <div>年龄: {{ item.age }}</div>
          <hr>
        </div>
      </div>
    </div>
  </div>
</div>
<script>
  const app = Vue.createApp({
    data() {
      return {
        dataLoaded: false, // 标记数据是否已加载
        items: [] // 存储数据的数组
      };
    },
    methods: {
      handleLoad() {
        const vm = this;
        // 发送GET请求获取数据
        $.ajax({
          url: 'http://127.0.0.1:5000/data',
          type: 'get',
          crossDomain: true, // 启用跨域请求
          dataType: 'json', // 指定数据类型为JSON
          success(data) {
            vm.items = data; // 将获取的数据存储到items数组中
            vm.dataLoaded = true; // 将数据加载状态设置为已加载
          }
        });
      },
      handleAdd() {
        const vm = this;
        const newData = {
          id: '123',
          name: 'New Data',
          age: 25
        };
        // 发送POST请求添加数据
        $.ajax({
          url: 'http://127.0.0.1:5000/data',
          type: 'post',
          data: JSON.stringify(newData),
          contentType: 'application/json',
          crossDomain: true, // 启用跨域请求
          dataType: 'json', // 指定数据类型为JSON
          success() {
            vm.handleLoad(); // 添加数据成功后重新加载数据
          }
        });
      },
      handleUpdate() {
        const vm = this;
        const updatedData = {
          name: 'Updated Data',
          age: 30
        };
        // 发送PUT请求更新数据
        $.ajax({
          url: 'http://127.0.0.1:5000/data/123',
          type: 'put',
          data: JSON.stringify(updatedData),
          contentType: 'application/json',
          crossDomain: true, // 启用跨域请求
          dataType: 'json', // 指定数据类型为JSON
          success() {
            vm.handleLoad(); // 更新数据成功后重新加载数据
          }
        });
      },
      handleDelete() {
        const vm = this;
        // 发送DELETE请求删除数据
        $.ajax({
          url: 'http://127.0.0.1:5000/data/123',
          type: 'delete',
          crossDomain: true, // 启用跨域请求
          success() {
            vm.handleLoad(); // 删除数据成功后重新加载数据
          }
        });
      }
    },
    mounted() {
      console.log('当前状态：mounted');
      this.handleLoad();
    }
  });

  app.mount('#app');
</script>
</body>

</html>
```

## 五、使用fetch前后端交互

更加简单的数据获取方式，功能更强大、更灵活，可以看做是的升级版

### 5.1 fetch基本用法

```js
fetch(url)
  .then(response => response.json())  // 解析响应为 JSON
  .then(data => console.log(data))    // 打印出解析后的数据,这里得到的才是真数据
  .catch(error => console.log('Error:', error)); // 捕获并打印任何出错信息
```

### 5.2 fetch请求参数**常用配置选项**

1. **method（String）**：定义 HTTP 请求方法。默认值为 'GET'。其他常见的值包括 'POST'、'PUT'、'DELETE'、'HEAD' 等。
2. **body（String / FormData / Blob / ArrayBufferView / ArrayBuffer / URLSearchParams / ReadableStream / FormData）**：定义请求的 body，用于 'POST' 或 'PUT' 请求。需要注意的是，你需要将 body 的数据转换为合适的格式，并在 headers 中设置正确的 'Content-Type'。
3. **headers（Object）**：定义请求的 HTTP 头。这是一个普通对象，其中的每个键值对都表示一个 HTTP 头的名称和值。默认值为空对象（{}）。

以下是一个具有这些配置选项的 fetch 请求示例：

```js
fetch('http://127.0.0.1:5000/data', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    key1: 'value1',
    key2: 'value2',
  }),
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.log('Error:', error));
```

在这个例子中，我们发送了一个 'POST' 请求到 'http://127.0.0.1:5000/data'。我们在 headers 中设置了 'Content-Type' 为 'application/json'，并且把一个 JavaScript 对象转换为 JSON 格式的字符串作为请求的 body。

除了以上这些，fetch API 还支持许多其他的配置选项，包括：

- **mode（String）**：定义请求的模式，比如 'cors'、'no-cors'、'same-origin' 或 'navigate'。这将影响跨站请求的行为。
- **credentials（String）**：定义是否应该在请求中包含凭证（如 cookies）。可能的值包括 'include'、'same-origin'、'omit'。
- **cache（String）**：定义请求的缓存模式。可能的值包括 'default'、'no-store'、'reload'、'no-cache'、'force-cache'、'only-if-cached'。
- **redirect（String）**：定义如何处理重定向。可能的值包括 'follow'、'error'、'manual'。
- **referrer（String）**：定义请求的 referrer。可能的值是任何有效的 URL。
- **referrerPolicy（String）**：定义请求的 referrer 策略。可能的值包括 'no-referrer'、'no-referrer-when-downgrade'、'origin'、'origin-when-cross-origin'、'same-origin'、'strict-origin'、'strict-origin-when-cross-origin'、'unsafe-url'。
- **integrity（String）**：一个包含请求的子资源完整性描述（SRI）的字符串。

### 5.3 Response 对象

fetch() 返回的 Response 对象表示服务器对 fetch 请求的响应。Response 对象有很多属性，比如 status（表示 HTTP 响应码）和 headers（表示响应头）。你还可以调用一些方法来获取响应体的内容，比如 text() 或 json()。

```js
fetch(url)
  .then(response => response.json())
  .then(data => console.log(data));
```

在上述示例中，我们首先调用 fetch() 发送请求。然后，当服务器返回响应时，我们调用 json() 方法来解析响应体中的 JSON 数据。最后，我们把解析后的数据打印到控制台。

### 5.4 错误处理

你可以使用 catch() 方法来捕获 fetch() 中可能出现的任何错误。

```js
fetch(url)
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.log('Error:', error));

```

### 5.5 异步处理

由于 fetch() 返回的是一个 Promise，所以你可以使用 async/await 语法来简化你的代码。

```js
async function fetchData() {
  try {
    let response = await fetch(url);
    let data = await response.json();
    console.log(data);
  } catch (error) {
    console.log('Error:', error);
  }
}

fetchData();
```

在这个示例中，我们创建了一个异步函数 fetchData()，在这个函数中，我们使用 await 关键字来等待 fetch() 和 json() 方法的结果。如果在这个过程中出现任何错误，我们就会捕获这个错误并打印到控制台。

### 5.6 fetch 增删改查页面

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <link href="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
  <script src="vue.js"></script>
</head>

<body>
<div class="container-fluid">
  <div class="row">
    <div class="col-md-6 col-md-offset-3" style="margin-top: 20px">
      <div id="app">
        <button @click="handleLoad" v-if="!dataLoaded">加载数据</button>
        <button @click="handleAdd">添加数据</button>
        <button @click="handleUpdate">更新数据</button>
        <button @click="handleDelete">删除数据</button>
        <hr>
        <div v-for="item in items" :key="item.id">
          <div>ID: {{ item.id }}</div>
          <div>名称: {{ item.name }}</div>
          <div>年龄: {{ item.age }}</div>
          <hr>
        </div>
      </div>
    </div>
  </div>
</div>
<script>
  const app = Vue.createApp({
    data() {
      return {
        dataLoaded: false, // 标记数据是否已加载
        items: [] // 存储数据的数组
      };
    },
    methods: {
      handleLoad() {
        const vm = this;
        // 发送GET请求获取数据
        fetch('http://127.0.0.1:5000/data')
                .then(response => response.json())
                .then(data => {
                  vm.items = data; // 将获取的数据存储到items数组中
                  vm.dataLoaded = true; // 将数据加载状态设置为已加载
                })
                .catch(error => console.error(error));
      },
      handleAdd() {
        const vm = this;
        const newData = {
          id: '123',
          name: 'New Data',
          age: 25
        };
        // 发送POST请求添加数据
        fetch('http://127.0.0.1:5000/data', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(newData)
        })
                .then(() => vm.handleLoad()) // 添加数据成功后重新加载数据
                .catch(error => console.error(error));
      },
      handleUpdate() {
        const vm = this;
        const updatedData = {
          name: 'Updated Data',
          age: 30
        };
        // 发送PUT请求更新数据
        fetch('http://127.0.0.1:5000/data/123', {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(updatedData)
        })
                .then(() => vm.handleLoad()) // 更新数据成功后重新加载数据
                .catch(error => console.error(error));
      },
      handleDelete() {
        const vm = this;
        // 发送DELETE请求删除数据
        fetch('http://127.0.0.1:5000/data/123', {
          method: 'DELETE'
        })
                .then(() => vm.handleLoad()) // 删除数据成功后重新加载数据
                .catch(error => console.error(error));
      }
    },
    mounted() {
      console.log('当前状态：mounted');
      this.handleLoad();
    }
  });

  app.mount('#app');
</script>
</body>

</html>

```

## 六、前后端交互之axios

Axios 是一个基于 Promise 的 HTTP 库，可以用在浏览器和 node.js 中。

### 6.0 安装

使用 npm:

```bash
npm install axios
```

使用 bower:

```ruby
bower install axios
```

使用 cdn:

```xml
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
```

### 6.1 基本用法:

**执行 GET 请求：**

```js
axios.get('/user?ID=12345') // 使用axios.get方法发送一个GET请求，请求的URL为'/user?ID=12345'
  .then(function (response) {
    // 这是请求成功后执行的代码块。response是服务器的响应。
    console.log(response);
  })
  .catch(function (error) {
    // 这是请求失败后执行的代码块。error是失败的错误信息。
    console.log(error);
  })
  .finally(function () {
    // 这是无论请求成功还是失败都会执行的代码块。
  });
```

**执行 POST 请求：**

```js
axios.get('/user?ID=12345') // 使用axios.get方法发送一个GET请求，请求的URL为'/user?ID=12345'
  .then(function (response) {
    // 这是请求成功后执行的代码块。response是服务器的响应。
    console.log(response);
  })
  .catch(function (error) {
    // 这是请求失败后执行的代码块。error是失败的错误信息。
    console.log(error);
  })
  .finally(function () {
    // 这是无论请求成功还是失败都会执行的代码块。
  });

```

### 6.2 **执行多个并发请求：**

你可以使用 `axios.all` 方法来发送多个并发请求：

```js
axios.all([
  axios.get('/user/12345'),
  axios.get('/user/67890')
])
.then(axios.spread(function (userResp, user2Resp) {
  // 两个请求现在都已经完成
  console.log(userResp.data);
  console.log(user2Resp.data);
}));
```

### 6.3 错误处理

使用 `catch` 方法来处理请求过程中的任何错误：

```js
axios.get('/user/12345')
  .then(function (response) {
    console.log(response.data);
  })
  .catch(function (error) {
    console.log(error);
  });

```

### 6.4 使用 async/await

如果你的环境支持 async/await，那么可以使用这些特性来简化你的代码：

```js
async function getUser() {
  try {
    const response = await axios.get('/user/12345');
    console.log(response.data);
  } catch (error) {
    console.log(error);
  }
}
getUser();
```

### 6.5 axios 返回的响应对象主要属性：

- **data**：服务器响应的数据。这是我们通常最关心的部分，它包含了服务器返回的数据。
- **headers**：服务器响应的头信息。这是一个对象，包含了响应的各种头部信息，如 'Content-Type'。
- **status**：服务器响应的 HTTP 状态码。比如，200 表示请求成功，404 表示未找到，500 表示服务器错误等。
- **statusText**：HTTP 状态消息。与状态码相对应，例如 "OK" 对应 200，"Not Found" 对应 404 等。

你可以使用 `.then` 方法获取这些信息，例如：

```javascript
axios.get('/api/data')
  .then(response => {
    console.log(response.data);  // 访问返回的数据
    console.log(response.status);  // 访问状态码
    console.log(response.statusText);  // 访问状态消息
    console.log(response.headers);  // 访问响应头
  })
  .catch(error => {
    console.log(error);
  });
```

### 6.6 axios 常用的配置选项

1. **url (String)**：要请求的服务器 URL。
2. **method (String)**：请求方法。默认是 'GET'。
3. **baseURL (String)**：将被添加到 `url` 前面，除非 `url` 是绝对的。
4. **headers (Object)**：自定义请求头。
5. **params (Object)**：要与请求一起发送的 URL 参数，必须是纯对象或 URLSearchParams 对象。
6. **data (Object, String, ArrayBuffer, ArrayBufferView, URLSearchParams, FormData, File, Blob)**：作为请求主体发送的数据。只适用于 'PUT', 'POST', 'PATCH' 方法。
7. **timeout (Number)**：指定请求超时之前的毫秒数。
8. **withCredentials (Boolean)**：表示跨站点访问控制请求是否应该使用证书，如 cookies, authorization headers 或 TLS client certificates。默认为 false。
9. **responseType (String)**：表示服务器将响应的数据类型。可能的值是 'arraybuffer', 'blob', 'document', 'json', 'text', 'stream'。
10. **onUploadProgress (Function)**：允许处理上传的进度事件。
11. **onDownloadProgress (Function)**：允许处理下载的进度事件。
12. **validateStatus (Function)**：定义是否将 promise 解析为成功或失败，返回 true（或设置为 null 或 undefined）将 promise 状态设置为成功，返回 false 将状态设置为失败。

```js
axios({
  url: '/user', // URL路径，将拼接到baseURL之后
  method: 'post', // 使用的HTTP方法
  baseURL: 'http://127.0.0.1:5000/', // 基础URL，除非URL是绝对路径，否则将被添加到URL之前
  headers: {'X-Requested-With': 'XMLHttpRequest'}, // 自定义请求头
  params: {
    ID: 12345 // URL查询参数，这将添加到URL后面 "?ID=12345"
  },
  data: {
    firstName: 'Fred', // 请求体数据，将转换为JSON字符串发送给服务器
    lastName: 'Flintstone'
  },
  timeout: 1000, // 请求超时时间，单位毫秒
  withCredentials: true, // 是否允许发送cookie，只有在同源请求时才会发送
  responseType: 'json', // 服务器响应的数据类型
})
.then(function(response) { // 请求成功的回调函数，response是服务器响应的信息
  console.log(response);
})
.catch(function(error) { // 请求失败的回调函数，error是失败的错误信息
  console.log(error);
});

```

这个例子中，我们向 'http://127.0.0.1:5000/' 发送了一个 POST 请求，带有 query 参数 ID=12345 和请求主体，同时设置了一些其他配置。

### 6.7 **拦截请求和响应**

Axios允许你在请求发送给服务器之前，或当服务器的响应返回到then或catch方法之前，拦截请求或响应。拦截器可以对请求和响应进行预处理，对于处理API的统一请求参数、响应错误等通用处理场景十分有用。

- 请求拦截器：请求拦截器主要用于在请求发送前修改请求配置，比如设置通用的请求头、设置请求超时时间、在请求中携带用户token等等。请求拦截器接收请求配置作为参数，并且预计返回修改后的配置。
- 响应拦截器：响应拦截器主要用于处理请求返回的结果。我们可以在这里对返回结果进行提前处理，如统一处理错误信息、对返回结果进行提前处理等等。响应拦截器接收响应作为参数，并且预计返回处理后的响应。

```js
// 添加请求拦截器
axios.interceptors.request.use(function (config) {
    // 在这里，我们可以修改请求配置，比如添加通用的请求头
    config.headers.common['Authorization'] = 'Bearer token';
    return config;  // 返回修改后的请求配置
}, function (error) {
    // 如果请求出错，我们在这里处理错误，例如显示错误提示
    return Promise.reject(error);
});

// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    // 在这里，我们可以处理返回的响应。例如，如果服务器返回的状态不是 'success'，我们可以抛出错误
    if (response.data.status !== 'success') {
        console.error('Data error:', response.data.message);
    }
    return response;  // 返回处理后的响应
}, function (error) {
    // 如果响应出错，我们在这里处理错误，例如显示错误提示
    return Promise.reject(error);
});

```

### 6.8 **取消请求**

有时，我们可能需要取消一个正在进行的HTTP请求。这可能是因为用户已经不再需要请求的结果，或者是因为我们需要避免并发请求。axios提供了一个取消令牌，我们可以使用它来取消请求。需要注意的是，一旦取消了请求，那么请求就无法再次使用。

```js
var CancelToken = axios.CancelToken;
var source = CancelToken.source();

axios.get('/user/12345', {
    cancelToken: source.token  // 添加取消令牌到请求
}).catch(function (thrown) {
    if (axios.isCancel(thrown)) {  // 如果请求被取消，axios.isCancel会返回true
        console.log('Request canceled', thrown.message);
    } else {
        // 处理其他类型的错误
    }
});

// 在需要的时候取消请求
source.cancel('Operation canceled by the user.');
```

### **6.9 转换 JSON 数据：**

默认情况下，axios会自动将请求或响应的数据转换为JSON。这意味着，当我们发送请求时，如果请求体中的数据是一个JavaScript对象，那么axios会自动将其转换为JSON字符串。同样，当我们接收到响应时，如果响应体中的数据是JSON，那么axios会自动将其转换为JavaScript对象。

```js
axios.get('/user/12345')
    .then(function (response) {
        console.log(response.data);  // response.data 已经被转换为 JavaScript 对象
    });
```

### **6.10 客户端防御 XSRF：**

XSRF，也被称为跨站请求伪造，是一种常见的网络攻击手段。为了防止XSRF攻击，axios提供了一种机制：当浏览器环境设置了一个标准的 XSRF-TOKEN cookie时，你可以配置 axios 自动将它添加到请求头中，这样服务器可以验证请求的合法性。

```js
// 如果浏览器环境设置了一个名为 'XSRF-TOKEN' 的 cookie，axios 会自动将它添加到名为 'X-XSRF-TOKEN' 的请求头中
axios.defaults.xsrfCookieName = 'XSRF-TOKEN';
axios.defaults.xsrfHeaderName = 'X-XSRF-TOKEN';
```

### 6.11  axios 增删改查 页面

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <link href="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
  <script src="vue.js"></script>
  <!-- 引入 Axios -->
  <script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
</head>

<body>
<div class="container-fluid">
  <div class="row">
    <div class="col-md-6 col-md-offset-3" style="margin-top: 20px">
      <div id="app">
        <button @click="handleLoad" v-if="!dataLoaded">加载数据</button>
        <button @click="handleAdd">添加数据</button>
        <button @click="handleUpdate">更新数据</button>
        <button @click="handleDelete">删除数据</button>
        <hr>
        <div v-for="item in items" :key="item.id">
          <div>ID: {{ item.id }}</div>
          <div>名称: {{ item.name }}</div>
          <div>年龄: {{ item.age }}</div>
          <hr>
        </div>
      </div>
    </div>
  </div>
</div>
<script>
  const app = Vue.createApp({
    data() {
      return {
        dataLoaded: false, // 标记数据是否已加载
        items: [] // 存储数据的数组
      };
    },
    methods: {
      handleLoad() {
        const vm = this;
        // 使用axios发送GET请求获取数据
        axios.get('http://127.0.0.1:5000/data')
                .then(response => {
                  vm.items = response.data; // 将获取的数据存储到items数组中
                  vm.dataLoaded = true; // 将数据加载状态设置为已加载
                })
                .catch(error => console.error(error));
      },
      handleAdd() {
        const vm = this;
        const newData = {
          id: '123',
          name: 'New Data',
          age: 25
        };
        // 使用axios发送POST请求添加数据
        axios.post('http://127.0.0.1:5000/data', newData)
                .then(() => vm.handleLoad()) // 添加数据成功后重新加载数据
                .catch(error => console.error(error));
      },
      handleUpdate() {
        const vm = this;
        const updatedData = {
          name: 'Updated Data',
          age: 30
        };
        // 使用axios发送PUT请求更新数据
        axios.put('http://127.0.0.1:5000/data/123', updatedData)
                .then(() => vm.handleLoad()) // 更新数据成功后重新加载数据
                .catch(error => console.error(error));
      },
      handleDelete() {
        const vm = this;
        // 使用axios发送DELETE请求删除数据
        axios.delete('http://127.0.0.1:5000/data/123')
                .then(() => vm.handleLoad()) // 删除数据成功后重新加载数据
                .catch(error => console.error(error));
      }
    },
    mounted() {
      console.log('当前状态：mounted');
      this.handleLoad();
    }
  });

  app.mount('#app');
</script>
</body>

</html>
```


