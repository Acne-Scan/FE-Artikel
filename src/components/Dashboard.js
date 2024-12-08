import { useEffect, useState, useCallback } from "react";
import axios from "axios";
import ArticleForm from "./ArticleForm";
import ArticleTable from "./ArticleTable";

const Dashboard = () => {

  const API_URL = process.env.REACT_APP_API_URL;
  const article_URL = `${API_URL}/articles`;
  const [articles, setArticles] = useState([]);
  const [editingArticle, setEditingArticle] = useState(null);

  // Menggunakan useCallback untuk mendeklarasikan fetchArticles agar tetap stabil
  const fetchArticles = useCallback(async () => {
    try {
      const response = await axios.get(article_URL);
      console.log(response.data);
      if (response.data && response.data.data) {
        setArticles(response.data.data);
      }
    } catch (error) {
      console.error("Error fetching articles:", error);
    }
  }, [article_URL]); // Menggunakan article_URL sebagai dependensi jika diperlukan

  // tambah/update artikel
  const handleSave = async (article) => {
    try {
      const formData = new FormData();
      formData.append("name", article.name);
      formData.append("description", article.description);
      formData.append("author", article.author);
      if (article.image) {
        formData.append("image", article.image); // Tambahkan file gambar
      }
      if (editingArticle) {
        await axios.put(`${API_URL}/${editingArticle.article_id}`, formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        setEditingArticle(null);
      } else {
        await axios.post(API_URL, formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });
      }
      await fetchArticles();
    } catch (error) {
      console.error("Error saving article:", error);
    }
  };

  // hapus artikel
  const handleDelete = async (id) => {
    try {
      await axios.delete(`${API_URL}/${id}`);
      await fetchArticles();
    } catch (error) {
      console.error("Error deleting article:", error);
    }
  };

  // edit artikel
  const handleEdit = (article) => {
    setEditingArticle(article);
  };

  useEffect(() => {
    fetch('https://acnescan-final.et.r.appspot.com/articles')
      .then(response => response.json())
      .then(data => console.log(data))
      .catch(error => console.log('Error:', error));
  }, []);

  return (
    <div className="container mt-5">
      <h1 className="mb-4 text-center">Dashboard</h1>
      <div className="row">
        <div className="col-md-5">
          <div className="p-4 border rounded mb-4">
            <h3 className="mb-3">
              {editingArticle ? "Edit Article" : "Add Article"}
            </h3>
            <ArticleForm onSave={handleSave} initialData={editingArticle} />
          </div>
        </div>
        <div className="col-md-7">
          <div className="p-4 border rounded">
            <h3 className="mb-3">Articles</h3>
            <ArticleTable
              articles={articles}
              onEdit={handleEdit}
              onDelete={handleDelete}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
