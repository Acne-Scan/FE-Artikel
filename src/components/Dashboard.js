import React, { useEffect, useState } from "react";
// import axios from "axios"; // aktifin pas backend ada
import ArticleForm from "./ArticleForm";
import ArticleTable from "./ArticleTable";

const Dashboard = () => {
  // const API_URL = "http://<prod-base-url>:8080/articles"; // URL API backend
  const [articles, setArticles] = useState([]);
  const [editingArticle, setEditingArticle] = useState(null);





// fetch data dari API
// const fetchArticles = async () => {
//   try {
//     const response = await axios.get(API_URL);
//     setArticles(response.data.data); // Sesuaikan dengan struktur response API
//   } catch (error) {
//     console.error("Error fetching articles:", error);
//   }
// };



//coba pakai yang atas ^ ini hanya testing work CRUDnya atau tidak

  // Dummy data untuk testing
  const fetchArticles = () => {
    setArticles([
      {
        id: 1,
        title: "Contoh Artikel",
        content: "Ini adalah isi artikel contoh.",
        author: "Penulis A",
        image: "https://via.placeholder.com/50",
      },
      {
        id: 2,
        title: "Artikel Kedua",
        content: "Isi artikel kedua di sini.",
        author: "Penulis B",
        image: "https://via.placeholder.com/50",
      },
    ]);
  };




// tambah/update artikel

  // const handleSave = async (article) => {
  //   try {
  //     const formData = new FormData();
  //     formData.append("name", article.title);
  //     formData.append("description", article.content);
  //     formData.append("author", article.author);
  //     if (article.image) {
  //       formData.append("image", article.image); // Tambahkan file gambar
  //     }
  //     if (editingArticle) {
  //       await axios.put(`${API_URL}/${editingArticle.id}`, formData, {
  //         headers: { "Content-Type": "multipart/form-data" },
  //       });
  //       setEditingArticle(null);
  //     } else {
  //       await axios.post(API_URL, formData, {
  //         headers: { "Content-Type": "multipart/form-data" },
  //       });
  //     }
  //     await fetchArticles();
  //   } catch (error) {
  //     console.error("Error saving article:", error);
  //   }
  // };


//coba pakai yang atas ^ ini hanya testing work CRUDnya atau tidak

    // Update artikel

const handleSave = (article) => {
  if (editingArticle) {

    const updatedArticles = articles.map((item) =>
      item.id === editingArticle.id
        ? {
            ...article,
            id: editingArticle.id,
            image: article.image
              ? URL.createObjectURL(article.image) // gambar baru
              : editingArticle.image, 
          }
        : item
    );
    setArticles(updatedArticles);
    setEditingArticle(null);
  } else {

    // nambah artikel baru
    setArticles([
      ...articles,
      {
        ...article,
        id: articles.length + 1,
        image: article.image
          ? URL.createObjectURL(article.image) 
          : "https://via.placeholder.com/50", 
      },
    ]);
  }
};






  // hapus artikel

  // const handleDelete = async (id) => {
  //   try {
  //     await axios.delete(`${API_URL}/${id}`);
  //     await fetchArticles();
  //   } catch (error) {
  //     console.error("Error deleting article:", error);
  //   }
  // };


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