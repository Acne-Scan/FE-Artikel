import React from "react";

const ArticleTable = ({ articles, onEdit, onDelete }) => {
  return (
    <table className="table table-striped">
      <thead>
        <tr>
          <th>Title</th>
          <th>Author</th>
          <th>Content</th>
          <th>Image</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {articles.length === 0 ? (
          <tr>
            <td colSpan="5" className="text-center">
              No articles available.
            </td>
          </tr>
        ) : (
          articles.map((article) => (
            <tr key={article.id}>
              <td>{article.title}</td>
              <td>{article.author}</td>
              <td>{article.content}</td>
              <td>
                <img src={article.image} alt={article.title} width="50" />
              </td>
              <td>
                <button
                  className="btn btn-warning btn-sm me-2"
                  onClick={() => onEdit(article)}
                >
                  Edit
                </button>
                <button
                  className="btn btn-danger btn-sm"
                  onClick={() => onDelete(article.id)}
                >
                  Delete
                </button>
              </td>
            </tr>
          ))
        )}
      </tbody>
    </table>
  );
};

export default ArticleTable;
