package apis

import (
	"github.com/gin-gonic/gin"
)

const INDEX = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>GO-ADMIN 接口测试</title>
<style>
body{
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 20px;
  background-color: #f5f5f5;
}
.container {
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}
h1 {
  color: #333;
  margin-bottom: 40px;
}
.button-group {
  display: flex;
  gap: 20px;
  justify-content: center;
  flex-wrap: wrap;
}
.button {
  padding: 15px 40px;
  font-size: 18px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  text-decoration: none;
  display: inline-block;
}
.button-swagger {
  background-color: #007bff;
  color: white;
}
.button-swagger:hover {
  background-color: #0056b3;
}
.button-test {
  background-color: #28a745;
  color: white;
}
.button-test:hover {
  background-color: #218838;
}
</style>
</head>
<body>
<div class="container">
  <h1>GO-ADMIN 接口测试平台</h1>
  <div class="button-group">
    <a href="/swagger/admin/index.html" class="button button-swagger">Swagger 文档</a>
    <a href="/test-page" class="button button-test">接口测试页面</a>
  </div>
</div>
</body>
</html>
`

func GoAdmin(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, INDEX)
}

const TEST_PAGE = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>GO-ADMIN 接口测试</title>
<style>
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 20px;
  background-color: #f5f5f5;
}
.container {
  max-width: 1200px;
  margin: 0 auto;
}
h1 {
  color: #333;
  text-align: center;
}
.section {
  background: white;
  padding: 20px;
  margin-bottom: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
h2 {
  color: #007bff;
  margin-top: 0;
}
.form-group {
  margin-bottom: 15px;
}
label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}
input, textarea, select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
}
button {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
button:hover {
  background-color: #0056b3;
}
.response {
  margin-top: 15px;
  padding: 15px;
  background-color: #f8f9fa;
  border-radius: 4px;
  white-space: pre-wrap;
  word-break: break-all;
}
.captcha-img {
  margin-top: 10px;
  max-width: 200px;
}
</style>
</head>
<body>
<div class="container">
  <h1>接口测试页面</h1>
  
  <!-- 登录相关接口 -->
  <div class="section">
    <h2>登陆接口</h2>
    
    <div class="form-group">
      <h3>获取验证码</h3>
      <button onclick="getCaptcha()">获取验证码</button>
      <div id="captcha-container"></div>
      <div class="response" id="captcha-response"></div>
    </div>
    
    <div class="form-group">
      <h3>用户登录</h3>
      <label>用户名</label>
      <input type="text" id="login-username" value="admin">
      <label>密码</label>
      <input type="password" id="login-password" value="123456">
      <label>验证码ID</label>
      <input type="text" id="login-captcha-id">
      <label>验证码</label>
      <input type="text" id="login-captcha">
      <button onclick="login()">登录</button>
      <div class="response" id="login-response"></div>
    </div>
    
    <div class="form-group">
      <h3>刷新Token</h3>
      <button onclick="refreshToken()">刷新Token</button>
      <div class="response" id="refresh-response"></div>
    </div>
  </div>
  
  <!-- 配置管理相关接口 -->
  <div class="section">
    <h2>配置管理接口</h2>
    
    <div class="form-group">
      <h3>获取系统前台配置（无需权限）</h3>
      <button onclick="getAppConfig()">获取配置</button>
      <div class="response" id="app-config-response"></div>
    </div>
    
    <div class="form-group">
      <h3>获取配置管理列表</h3>
      <label>配置名称</label>
      <input type="text" id="config-name">
      <label>配置Key</label>
      <input type="text" id="config-key">
      <label>页码</label>
      <input type="number" id="config-page-index" value="1">
      <label>每页条数</label>
      <input type="number" id="config-page-size" value="10">
      <button onclick="getConfigList()">获取列表</button>
      <div class="response" id="config-list-response"></div>
    </div>
    
    <div class="form-group">
      <h3>获取单条配置</h3>
      <label>配置ID</label>
      <input type="number" id="config-id">
      <button onclick="getConfig()">获取配置</button>
      <div class="response" id="config-response"></div>
    </div>
    
    <div class="form-group">
      <h3>创建配置</h3>
      <label>配置名称</label>
      <input type="text" id="create-config-name">
      <label>配置Key</label>
      <input type="text" id="create-config-key">
      <label>配置值</label>
      <input type="text" id="create-config-value">
      <label>配置类型</label>
      <select id="create-config-type">
        <option value="Y">是</option>
        <option value="N">否</option>
      </select>
      <button onclick="createConfig()">创建配置</button>
      <div class="response" id="create-config-response"></div>
    </div>
    
    <div class="form-group">
      <h3>修改配置</h3>
      <label>配置ID</label>
      <input type="number" id="update-config-id">
      <label>配置名称</label>
      <input type="text" id="update-config-name">
      <label>配置Key</label>
      <input type="text" id="update-config-key">
      <label>配置值</label>
      <input type="text" id="update-config-value">
      <button onclick="updateConfig()">修改配置</button>
      <div class="response" id="update-config-response"></div>
    </div>
    
    <div class="form-group">
      <h3>删除配置</h3>
      <label>配置ID数组（如: [1,2,3]）</label>
      <input type="text" id="delete-config-ids" value="[1]">
      <button onclick="deleteConfig()">删除配置</button>
      <div class="response" id="delete-config-response"></div>
    </div>
    
    <div class="form-group">
      <h3>获取设置配置</h3>
      <button onclick="getSetConfig()">获取设置配置</button>
      <div class="response" id="get-set-config-response"></div>
    </div>
    
    <div class="form-group">
      <h3>设置配置</h3>
      <label>配置数据（JSON数组）</label>
      <textarea id="set-config-data" rows="4">[{"configKey":"test","configValue":"value"}]</textarea>
      <button onclick="setConfig()">设置配置</button>
      <div class="response" id="set-config-response"></div>
    </div>
  </div>
</div>

<script>
let token = '';
let captchaId = '';

function setAuthHeader(xhr) {
  if (token) {
    xhr.setRequestHeader('Authorization', 'Bearer ' + token);
  }
}

function displayResponse(elementId, data, status = 'success') {
  const element = document.getElementById(elementId);
  element.style.color = status === 'success' ? '#28a745' : '#dc3545';
  element.textContent = JSON.stringify(data, null, 2);
}

function getCaptcha() {
  fetch('/api/v1/captcha')
    .then(response => response.json())
    .then(data => {
      captchaId = data.id;
      document.getElementById('login-captcha-id').value = captchaId;
      const container = document.getElementById('captcha-container');
      container.innerHTML = '<img src="' + data.data + '" alt="验证码" class="captcha-img">';
      displayResponse('captcha-response', data);
    })
    .catch(error => displayResponse('captcha-response', {error: error.message}, 'error'));
}

function login() {
  const username = document.getElementById('login-username').value;
  const password = document.getElementById('login-password').value;
  const captchaIdVal = document.getElementById('login-captcha-id').value;
  const captchaVal = document.getElementById('login-captcha').value;
  
  fetch('/api/v1/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      username: username,
      password: password,
      captchaId: captchaIdVal,
      captcha: captchaVal
    })
  })
  .then(response => response.json())
  .then(data => {
    if (data.code === 200 && data.data && data.data.token) {
      token = data.data.token;
    }
    displayResponse('login-response', data);
  })
  .catch(error => displayResponse('login-response', {error: error.message}, 'error'));
}

function refreshToken() {
  const xhr = new XMLHttpRequest();
  xhr.open('GET', '/api/v1/refresh_token');
  setAuthHeader(xhr);
  xhr.onload = function() {
    const data = JSON.parse(xhr.responseText);
    if (data.code === 200 && data.data && data.data.token) {
      token = data.data.token;
    }
    displayResponse('refresh-response', data, xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('refresh-response', {error: '请求失败'}, 'error');
  };
  xhr.send();
}

function getAppConfig() {
  fetch('/api/v1/app-config')
    .then(response => response.json())
    .then(data => displayResponse('app-config-response', data))
    .catch(error => displayResponse('app-config-response', {error: error.message}, 'error'));
}

function getConfigList() {
  const name = document.getElementById('config-name').value;
  const key = document.getElementById('config-key').value;
  const pageIndex = document.getElementById('config-page-index').value;
  const pageSize = document.getElementById('config-page-size').value;
  
  let url = '/api/v1/sys-config?pageIndex=' + pageIndex + '&pageSize=' + pageSize;
  if (name) url += '&configName=' + encodeURIComponent(name);
  if (key) url += '&configKey=' + encodeURIComponent(key);
  
  const xhr = new XMLHttpRequest();
  xhr.open('GET', url);
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('config-list-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('config-list-response', {error: '请求失败'}, 'error');
  };
  xhr.send();
}

function getConfig() {
  const id = document.getElementById('config-id').value;
  if (!id) {
    displayResponse('config-response', {error: '请输入配置ID'}, 'error');
    return;
  }
  
  const xhr = new XMLHttpRequest();
  xhr.open('GET', '/api/v1/sys-config/' + id);
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('config-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('config-response', {error: '请求失败'}, 'error');
  };
  xhr.send();
}

function createConfig() {
  const name = document.getElementById('create-config-name').value;
  const key = document.getElementById('create-config-key').value;
  const value = document.getElementById('create-config-value').value;
  const type = document.getElementById('create-config-type').value;
  
  const xhr = new XMLHttpRequest();
  xhr.open('POST', '/api/v1/sys-config');
  xhr.setRequestHeader('Content-Type', 'application/json');
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('create-config-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('create-config-response', {error: '请求失败'}, 'error');
  };
  xhr.send(JSON.stringify({
    configName: name,
    configKey: key,
    configValue: value,
    configType: type
  }));
}

function updateConfig() {
  const id = document.getElementById('update-config-id').value;
  if (!id) {
    displayResponse('update-config-response', {error: '请输入配置ID'}, 'error');
    return;
  }
  
  const name = document.getElementById('update-config-name').value;
  const key = document.getElementById('update-config-key').value;
  const value = document.getElementById('update-config-value').value;
  
  const xhr = new XMLHttpRequest();
  xhr.open('PUT', '/api/v1/sys-config/' + id);
  xhr.setRequestHeader('Content-Type', 'application/json');
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('update-config-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('update-config-response', {error: '请求失败'}, 'error');
  };
  xhr.send(JSON.stringify({
    id: parseInt(id),
    configName: name,
    configKey: key,
    configValue: value
  }));
}

function deleteConfig() {
  const idsStr = document.getElementById('delete-config-ids').value;
  let ids;
  try {
    ids = JSON.parse(idsStr);
  } catch (e) {
    displayResponse('delete-config-response', {error: '无效的ID数组格式'}, 'error');
    return;
  }
  
  const xhr = new XMLHttpRequest();
  xhr.open('DELETE', '/api/v1/sys-config');
  xhr.setRequestHeader('Content-Type', 'application/json');
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('delete-config-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('delete-config-response', {error: '请求失败'}, 'error');
  };
  xhr.send(JSON.stringify(ids));
}

function getSetConfig() {
  const xhr = new XMLHttpRequest();
  xhr.open('GET', '/api/v1/set-config');
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('get-set-config-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('get-set-config-response', {error: '请求失败'}, 'error');
  };
  xhr.send();
}

function setConfig() {
  const dataStr = document.getElementById('set-config-data').value;
  let data;
  try {
    data = JSON.parse(dataStr);
  } catch (e) {
    displayResponse('set-config-response', {error: '无效的JSON格式'}, 'error');
    return;
  }
  
  const xhr = new XMLHttpRequest();
  xhr.open('PUT', '/api/v1/set-config');
  xhr.setRequestHeader('Content-Type', 'application/json');
  setAuthHeader(xhr);
  xhr.onload = function() {
    displayResponse('set-config-response', JSON.parse(xhr.responseText), xhr.status === 200 ? 'success' : 'error');
  };
  xhr.onerror = function() {
    displayResponse('set-config-response', {error: '请求失败'}, 'error');
  };
  xhr.send(JSON.stringify(data));
}
</script>
</body>
</html>
`

func TestPage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, TEST_PAGE)
}

