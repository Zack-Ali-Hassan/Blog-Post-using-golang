import React, { useEffect } from "react";

const Home = ({ posts, fetchPosts }) => {
  useEffect(() => {
    // Fetch posts when the component mounts
    const fetchedPosts = fetchPosts();
    // Here we would normally set the posts state, but we're passing it down
  }, [fetchPosts]);

  return (
    <div className="container mt-5">
      <h2>Blog Posts</h2>
      {posts.length === 0 ? (
        <p>No posts available.</p>
      ) : (
        <div className="row mt-5">
          {posts.map((post) => (
            <div key={post.id} className="col-md-4 mb-4">
              <div className="card shadow">
                <div className="card-body">
                  <h5 className="card-title">{post.title}</h5>
                  <p className="card-text">{post.content}</p>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Home;
