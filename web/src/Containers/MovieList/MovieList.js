import React, { Component } from 'react';

export default class MovieList extends Component {
    componentWillMount() {
        const { fetchMovies } = this.props

        fetchMovies()
    }

    render() {
        const { loading } = this.props.movies

        return (
            <div>
                {loading && "LOADING"}
                {!loading && "Movie list"}
            </div>
        )
    }
}
