// src/components/BlogDetail.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import './ArticleDetail.css'; // Import CSS

const ArticleDetail = () => {
  const { id } = useParams();
  const [Article, setArticle] = useState(null);

  useEffect(() => {
    axios.get(`http://localhost:8000/api/vnexpress/${id}`)
      .then(response => {
        setArticle(response.data.record);
      })
      .catch(error => {
        console.error("Error!", error);
      });
  }, [id]);

  return (
    <div className="article-detail">
      {Article ? (
        <div>
          <h1>{Article.title}</h1>
          <img src={Article.image} alt={Article.title} />
          <p>{Article.description}</p>
          <a href={Article.link} target="_blank" rel="noopener noreferrer">Read more</a>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default ArticleDetail;
