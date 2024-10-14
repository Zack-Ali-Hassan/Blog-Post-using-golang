import React, { useEffect, useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import CreatePost from "./components/CreatePost";
import Home from "./components/Home";
import Navbar from "./components/NavBar";
import axios from "axios";
export const BASE_URL = import.meta.env.MODE == "development" ? "http://127.0.0.1:4444/api" : "/api";
function App() {
  const [posts, setPosts] = useState([]);
  const getPosts = async ()=>{
    const {data} = await axios.get(BASE_URL+"/post/")
    setPosts(data)
  }
  useEffect(() => {
   
    getPosts()
  }, []);


  return (
    <Router>
      <div>
        <Navbar />
        <Routes>
          <Route
            path="/"
            element={<Home posts={posts} getPosts ={getPosts} />}
          />
          <Route
            path="/create-post"
            element={<CreatePost getPosts ={getPosts} />}
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
