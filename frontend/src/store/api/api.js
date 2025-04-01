import axios from "axios";

const url = 'http://localhost:8080';

const sendGet = async (endpoint) => {
    try {
        const response = await axios.get(`${url}${endpoint}`);
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('Ошибка:', error);
    }
};

const sendPrivateGet = async (endpoint) => {
    try {
        const accessToken = localStorage.getItem('accessToken');
        const response = await axios.get(`${url}${endpoint}`, {
            headers: { Authorization: `Bearer ${accessToken}` },
        });
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('Ошибка:', error);
    }
};

const sendPost = async (endpoint, body) => {
    try {
        const response = await axios.post(`${url}${endpoint}`, body);
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('Ошибка:', `ошибка при пост запросе: ${error}`);
    }
};

const sendPrivatePost = async (endpoint, body) => {
    try {
        const accessToken = localStorage.getItem('accessToken');
        const response = await axios.post(`${url}${endpoint}`, body, {
            headers: {
                Authorization: `Bearer ${accessToken}`
            }
        });
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('Ошибка:', `ошибка при пост запросе: ${error}`);
    }
};

export const registerRequest = async (email, login, pass) => {
    const data = {
        email: email,
        username: login,
        password: pass
    };
    return sendPost('/api/v1/auth/signup', data);
};

export const loginRequest = async (email, pass) => {
    const data = {
        email: email,
        password: pass
    };
    return sendPost(`/api/v1/auth/login`, data);
};

export const getUserInfoRequest = async (accessToken) => {
    return sendPrivateGet('/api/v1/auth/me', accessToken);
};

export const logoutRequest = async () => {
    const data = {
        refreshToken: localStorage.getItem('refreshToken')
    };
    return sendPrivatePost('/api/v1/auth/logout', data);
};