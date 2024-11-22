import React, { useEffect, useState } from "react";
//import axios from "axios"; aktifin pas URL APInya udh ada
import ArticleForm from "./ArticleForm";
import ArticleTable from "./ArticleTable";

const Dashboard = () => {
  const API_URL = "http://localhost:4000/api/articles"; // URL API
  const [articles, setArticles] = useState([]);
  const [editingArticle, setEditingArticle] = useState(null);



// fetch data dari API
//   const fetchArticles = async () => {
//     try {
//       const response = await axios.get(API_URL);
//       setArticles(response.data);
//     } catch (error) {
//       console.error("Error fetching articles:", error);
//     }
//   };



//coba pakai yang atas ^ ini hanya testing work CRUDnya atau tidak
const fetchArticles = () => {
    setArticles([
      {
        id: 1,
        title: "Contoh Artikel",
        content: "Ini adalah isi artikel contoh.",
        author: "Penulis A",
      },
      {
        id: 2,
        title: "Artikel Kedua",
        content: "Isi artikel kedua di sini.",
        author: "Penulis B",
      },
    ]);
  };




// tambah/update artikel
//   const handleSave = async (article) => {
//     try {
//       if (editingArticle) {
//         // Update artikel
//         await axios.put(`${API_URL}/${editingArticle.id}`, article);
//         setEditingArticle(null);
//       } else {
//         // Tambah artikel baru
//         await axios.post(API_URL, article); // Tambahkan artikel ke API
//       }
//       await fetchArticles(); // Refresh data setelah operasi
//     } catch (error) {
//       console.error("Error saving article:", error);
//     }
//   };


//coba pakai yang atas ^ ini hanya testing work CRUDnya atau tidak
const handleSave = (article) => {
    if (editingArticle) {
      // Update artikel yang sudah ada
      const updatedArticles = articles.map((item) =>
        item.id === editingArticle.id ? { ...article, id: editingArticle.id } : item
      );
      setArticles(updatedArticles);
      setEditingArticle(null);
    } else {
      // Tambah artikel baru
      setArticles([...articles, { ...article, id: articles.length + 1 }]);
    }
  };




  // hapus artikel
//   const handleDelete = async (id) => {
//     try {
//       await axios.delete(`${API_URL}/${id}`);
//       fetchArticles(); // refresh data
//     } catch (error) {
//       console.error("Error deleting article:", error);
//     }
//   };


//coba pakai yang atas ^ ini hanya testing work CRUDnya atau tidak
const handleDelete = (id) => {
    const updatedArticles = articles.filter((item) => item.id !== id);
    setArticles(updatedArticles);
  };


  // edit artikel
  const handleEdit = (article) => {
    setEditingArticle(article);
  };

  useEffect(() => {
    fetchArticles();
  }, []);

  return (
    <div className="container mt-5">
      <h1 className="mb-4 text-center">Dashboard</h1>
      <div className="row">
        <div className="col-md-5">
          <div className="p-4 border rounded mb-4">
            <h3 className="mb-3">{editingArticle ? "Edit Article" : "Add Article"}</h3>
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