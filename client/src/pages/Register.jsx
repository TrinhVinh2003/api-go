import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function Register() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate(); // Khởi tạo useNavigate

  const submit = async (e) => {
    e.preventDefault(); // Ngăn chặn hành động mặc định của form
    try {
      const response = await fetch("http://localhost:8000/register", {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password }),
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
          <h1 className='h3 mb-3 fw-normal'>Register</h1>
          <div className="form-floating">
            <input
              type="text"
              className="form-control"
              id="floatingInput"
              required
              placeholder="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <label htmlFor="floatingInput">Name</label>
          </div>
          <div className="form-floating">
            <input
              type="email"
              className="form-control"
              id="floatingEmail"
              required
              placeholder="name@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <label htmlFor="floatingEmail">Email address</label>
          </div>
          <div className="form-floating">
            <input
              type="password"
              className="form-control"
              id="floatingPassword"
              required
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <label htmlFor="floatingPassword">Password</label>
          </div>
          <button className="btn btn-primary w-100 py-2" type="submit">Register</button>
        </form>
      </main>
    </div>
  );
}
