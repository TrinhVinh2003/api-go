import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import './ArticleList.css'; // Import CSS

const ArticleList = () => {
  const [articles, setArticles] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8000/api/vnexpress')
      .then(response => {
        console.log("API Response:", response.data); // Log response to debug
        setArticles(response.data.records || []); // Fallback to empty array if undefined
      })
      .catch(error => {
        console.error("Error fetching articles:", error);
        setArticles([]); // Ensure articles is an empty array on error
      });
  }, []);

  return (
    <div className="article-list">
      <h1>vnexpress</h1>
      <ul>
        {articles.length > 0 ? (
          articles.map(article => (
            <li key={article.id}>
              <Link to={`/articles/${article.id}`} className="article-link">
                <div className="img1">
                  <img src={article.image} alt={article.title} />
                </div>
                <div className="title">
                  <h2>{article.title}</h2>
                  <p>{article.description}</p>
                </div>
              </Link>
            </li>
          ))
        ) : (
          <p>No articles available.</p> // Message if no articles are available
        )}
      </ul>
    </div>
  );
};

export default ArticleList;
