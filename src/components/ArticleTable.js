import React from "react";

const ArticleTable = ({ articles, onEdit, onDelete }) => {
  return (
    <table className="table table-striped">
      <thead>
        <tr>
          <th>Title</th>
          <th>Description</th>
          <th>Image</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
      {articles.length === 0 ? (
        <tr>
          <td colSpan="4" className="text-center">
            No articles available.
          </td>
        </tr>
      ) : (
        articles.map((article) => (
          <tr key={article.article_id}>
            <td>{article.name}</td>
            <td>{article.description}</td>
            <td>
              <img
                src={article.image}
                alt={article.name}
                style={{ width: "50px", height: "50px", objectFit: "cover" }}
              />
            </td>
            <td>
              <button
                className="btn btn-sm btn-warning me-2"
                onClick={() => onEdit(article)}
              >
                Edit
              </button>
              <button
                className="btn btn-sm btn-danger"
                onClick={() => onDelete(article.article_id)}
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
