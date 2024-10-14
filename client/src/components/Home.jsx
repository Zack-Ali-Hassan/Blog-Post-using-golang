import axios from "axios";
import React, { useState } from "react";
import toast from "react-hot-toast";
import { BASE_URL } from "../App";

const Home = ({ posts, getPosts }) => {
  const [selectedPost, setSelectedPost] = useState(null); // For storing the post to be updated
  const [updatedTitle, setUpdatedTitle] = useState(""); // For updating the title
  const [updatedContent, setUpdatedContent] = useState(""); // For updating the content
  const handleUpdateClick = (post) => {
    setSelectedPost(post); // Set the post to be updated
    setUpdatedTitle(post.title); // Prefill the form with the existing title
    setUpdatedContent(post.content); // Prefill the form with the existing content
    const updateModal = new bootstrap.Modal(
      document.getElementById("updateModal")
    ); // Show modal
    updateModal.show();
  };

  const handleUpdateSubmit = (e) => {
    e.preventDefault();
    // Here you would send the updated post to the server, but for now, just log it
    console.log("Updated Post:", {
      id: selectedPost.id,
      title: updatedTitle,
      content: updatedContent,
    });
    // Close the modal after updating
    const updateModal = bootstrap.Modal.getInstance(
      document.getElementById("updateModal")
    );
    updateModal.hide();
  };
  const handleUpdate = (postId) => {
    if (confirm(`Are you sure you want to update this post id: ${postId}`)) {
      console.log(`Update post with ID: ${postId}`);
    }
  };

  const handleDelete = (postId) => {
    if (confirm(`Are you sure you want to update this post id: ${postId}`)) {
      try {
        const deletePost = async () => {
          const { result } = await axios.delete(BASE_URL + `/post/${postId}`);
          toast.success("Deleted post successfully");
          console.log(result);
        };
        deletePost();
        getPosts();
      } catch (error) {
        toast.error("Error deleting post....");
        console.log("Error from frontend in deleting post: ", error);
      }
      console.log(`Delete post with ID: ${postId}`);
    }
  };
  return (
    <div className="container mt-5">
      <h2>Blog Posts</h2>
      {posts.length === 0 ? (
        <p>No posts available.</p>
      ) : (
        <div className="row mt-5">
          {posts.map((post) => (
            <div key={post._id} className="col-md-4 mb-4">
              <div className="card shadow">
                <div className="card-body">
                  <h5 className="card-title">{post.title}</h5>
                  <p className="card-text">{post.content}</p>
                  <div className="d-flex justify-content-end">
                    <button
                      className="btn btn-outline-primary btn-sm me-2"
                      onClick={() => handleUpdateClick(post)}
                    >
                      <i className="bi bi-pencil-fill"></i> Update
                    </button>
                    <button
                      className="btn btn-outline-danger btn-sm"
                      onClick={() => handleDelete(post._id)}
                    >
                      <i className="bi bi-trash-fill"></i> Delete
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
      {/* Update Modal */}
      <div
        className="modal fade"
        id="updateModal"
        tabIndex="-1"
        aria-labelledby="updateModalLabel"
        aria-hidden="true"
      >
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-header">
              <h5 className="modal-title" id="updateModalLabel">
                Update Blog Post
              </h5>
              <button
                type="button"
                className="btn-close"
                data-bs-dismiss="modal"
                aria-label="Close"
              ></button>
            </div>
            <div className="modal-body">
              <form onSubmit={handleUpdateSubmit}>
                <div className="mb-3">
                  <label htmlFor="postTitle" className="form-label">
                    Title:
                  </label>
                  <input
                    type="text"
                    className="form-control"
                    id="postTitle"
                    value={updatedTitle}
                    onChange={(e) => setUpdatedTitle(e.target.value)}
                    required
                  />
                </div>
                <div className="mb-3">
                  <label htmlFor="postContent" className="form-label">
                    Content:
                  </label>
                  <textarea
                    className="form-control"
                    id="postContent"
                    value={updatedContent}
                    onChange={(e) => setUpdatedContent(e.target.value)}
                    required
                  ></textarea>
                </div>
                <button type="submit" className="btn btn-primary">
                  Edit info
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
