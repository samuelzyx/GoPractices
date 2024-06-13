import React from 'react';
import './Main.css';
import reel from './Moviereel.gif'
import { Link, useLocation } from 'react-router-dom';

const MovieDashboard = () => {
    const location = useLocation();

    return (
        <div>
            <section className={"Background"}>
                <header>
                    <div className={'Button-layout'}>
                        <Link to="/movies" className={location.pathname === '/movies' ? 'active' : ''}>
                            <button>Movies</button>
                        </Link>
                        <button>Coming soon!</button>
                        <button>Trending</button>
                        <button>Leaving Soon</button>
                        <Link to="/health" className={location.pathname === '/health' ? 'active' : ''}>
                            <button>Health</button>
                        </Link>
                    </div>
                    <h1 className={'Header'}>The Home Movie Depot</h1>
                </header>
                <img src={reel} className={'Movie-gif'} alt="logo"/>
            </section>
        </div>

    );

}

export default MovieDashboard;

