import React, { useState, useEffect } from "react";

const ArticleForm = ({ onSave, initialData }) => {
  const [formData, setFormData] = useState({
    title: "",
    content: "",
    author: "",
  });

  useEffect(() => {
    if (initialData) {
      setFormData(initialData);
    }
  }, [initialData]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave(formData);
    setFormData({ title: "", content: "", author: "" });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-3">
        <label className="form-label">Title</label>
        <input
          type="text"
          className="form-control"
          name="title"
          value={formData.title}
          onChange={handleChange}
          required
        />
      </div>
      <div className="mb-3">
        <label className="form-label">Content</label>
        <textarea
          className="form-control"
          name="content"
          rows="4"
          value={formData.content}
          onChange={handleChange}
          required
        ></textarea>
      </div>
      <div className="mb-3">
        <label className="form-label">Author</label>
        <input
          type="text"
          className="form-control"
          name="author"
          value={formData.author}
          onChange={handleChange}
          required
        />
      </div>
      <button type="submit" className="btn btn-primary">
        {initialData ? "Update" : "Add"}
      </button>
    </form>
  );
};

export default ArticleForm;
