import React from 'react';
import './Header.css'
import {useNavigate } from 'react-router-dom';

export default function HeaderComponent() {
const Navigator  = useNavigate();
function Login(){
  Navigator("/Login")
}
function Register(){
  Navigator("/register")
}
  return (
    <header>
    <nav className='navbar navbar-dark bg-dark'>
      <a className='navbar-brand' href="#">Fresher Manage System</a>
  
    
        <ul className="navbar-nav me-3 mb-2 mb-md-0">
          <li className="nav-item">
            <a className="nav-link " aria-current="page" onClick={Register}>Register</a>
          </li>
          <li className="nav-item">
            <a className="nav-link" onClick={Login}>Log in</a>
          </li>
        </ul>
   
    </nav>
  </header>
  )
}
