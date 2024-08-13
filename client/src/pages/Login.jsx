import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();
  const submit = async (e) => {
    e.preventDefault(); // Ngăn chặn hành động mặc định của form
    try {
      const response = await fetch("http://127.0.0.1:8000/login", {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({ email, password }),
      });
      if (response.ok) {
        // Nếu đăng ký thành công, điều hướng đến trang đăng nhập
        navigate('/login');
      } else {
        // Xử lý lỗi nếu có
        console.error('Registration failed:', response.statusText);
      }
    } catch (error) {
      console.error('Error during registration:', error);
    }
  };
  return (
    <div>
      <main className="form-signin w-100 m-auto">
        <form onSubmit={submit}>
           
            <div className="form-floating">
                <input type="email" className="form-control" id="floatingInput" placeholder="name@example.com" onChange={(e) => setEmail(e.target.value)}/>
                <label for="floatingInput">Email address</label>
            </div>
            <div className="form-floating">
                <input type="password" className="form-control" id="floatingPassword" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
                <label for="floatingPassword">Password</label>
            </div>

            
            <button className="btn btn-primary w-100 py-2" type="submit">Sign in</button>
         
        </form>
        </main>
    </div>
  )
}

export default Login
