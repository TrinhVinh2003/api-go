import React from 'react';
import Login from './pages/Login';
import { BrowserRouter, Routes,Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'
// import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
// import ArticleList from './components/ArticleList';
// import ArticleDetail from './components/ArticleDetail';
import './App.css';
import Register from './pages/Register';
import ArticleList from './components/ArticleList';
import HeaderComponent from './components/HeaderComponent';
import ArticleDetail from './components/ArticleDetail';
const App = () => {
  return (
    <BrowserRouter>
    < HeaderComponent />
      <Routes>
        <Route path='/login' element = {<Login/>} ></Route>
        <Route path='/register' element  = {<Register/>}></Route>
        <Route path="/" element={<ArticleList />} />
        <Route path="/articles/:id" element={<ArticleDetail />} />
      </Routes>
    
    </BrowserRouter>
    


      //   <Router>
      //   <Routes>
        
      //     <Route path="/" element={<ArticleList />} />
      //     <Route path="/articles/:id" element={<ArticleDetail />} />
      //   </Routes>
      // </Router>

  );
};

export default App;