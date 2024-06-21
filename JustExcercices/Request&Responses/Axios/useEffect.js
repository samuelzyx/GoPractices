import React, { useState, useEffect } from 'react';
import './Movies.css';
import axios from 'axios';

function Movies() {
  const [movies, setMovies] = useState([]);
  const [isLoading, setIsLoading] = useState(false); // Added state for loading indicator

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true); // Set loading indicator to true
      const response = await axios.get('http://localhost:8080/api/movies');
      setMovies(response.data);
      setIsLoading(false); // Set loading indicator to false
    };

    fetchData();
  }, []);

  return (
    <div className="movie-container">
      <h2 className="movie-heading">Movies</h2>
      {isLoading ? (
        <p className="loading-message">Loading movies...</p>
      ) : (
        <ul className="movie-list">
          {movies.map((movie) => (
            <li key={movie.title} className="movie-item">
              <h3 className="movie-title">{movie.title}</h3>
              <p className="movie-average-rating">
                Average Rating: {movie.average_rating}
              </p>
              <p className="movie-plot">Plot: {movie.Plot}</p>
              {/* You can display other movie details here based on your needs */}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Movies;
