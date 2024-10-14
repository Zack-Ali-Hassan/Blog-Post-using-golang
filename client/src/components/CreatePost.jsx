import axios from 'axios';
import React, { useState } from 'react';
import toast from 'react-hot-toast';
import { BASE_URL } from '../App';

const CreatePost = ({getPosts}) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    
    const createPost = async ()=>{
      try {
        const {result} = await axios.post(BASE_URL + "/post",{title : title, content: content})
        toast.success("Inserted successfully")
        console.log("Creating Post result are: ", result)
        getPosts()
      } catch (error) {
        console.log("Error from frontend in creating post: ", error);
        toast.error("Error.....")
      }
     
    }
    createPost()
    setTitle('');
    setContent('');
  };

  return (
    <div className="container mt-5">
      <h1>Create a New Post</h1>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">Title:</label>
          <input
            type="text"
            className="form-control"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label">Content:</label>
          <textarea
            className="form-control"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">Create Post</button>
      </form>
    </div>
  );
};

export default CreatePost;
