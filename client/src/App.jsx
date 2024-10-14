import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import CreatePost from './components/CreatePost';
import Home from './components/Home';
import Navbar from './components/NavBar';

function App() {
  const [posts, setPosts] = useState([]);

  // Simulated fetch function to get posts
  const fetchPosts = () => {
    // This is where you'd normally fetch from your backend
    return [
      { id: 1, title: "First Blog Post", content: "This is the content of the first blog post." },
      { id: 2, title: "Second Blog Post", content: "This is the content of the second blog post." },
      // Add more mock posts as needed
    ];
  };

  // Simulated function to add a post
  const addPost = (newPost) => {
    setPosts((prevPosts) => [...prevPosts, newPost]);
  };

  return (
    <Router>
      <div>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home posts={posts} fetchPosts={fetchPosts} />} />
          <Route path="/create-post" element={<CreatePost addPost={addPost} />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
