import axios from "axios";

const url = 'http://localhost:8080';

const sendGet = async (endpoint) => {
    try {
        const response = await axios.get(`${url}${endpoint}`);
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('Ошибка GET:', error);
    }
};

const sendPrivateGet = async (endpoint) => {
    try {
        const accessToken = localStorage.getItem('accessToken');
        let response = await axios.get(`${url}${endpoint}`, {
            headers: { Authorization: `Bearer ${accessToken}` },
        });
        console.log(response.data);

        if(response.status === 401) {
            const refreshResult = await refreshAccessTokenRequest(localStorage.getItem('refreshToken'));
            if(refreshResult) {
                response = await axios.get(`${url}${endpoint}`, {
                    headers: { Authorization: `Bearer ${accessToken}` },
                });
            }
            else {
                return null;
            }
        }

        return response.data;
    } catch (error) {
        console.error('Ошибка GET:', error);
    }
};

const sendPost = async (endpoint, body) => {
    try {
        const response = await axios.post(`${url}${endpoint}`, body);
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('Ошибка POST:', error);
    }
};

const sendPrivatePost = async (endpoint, body) => {
    try {
        const accessToken = localStorage.getItem('accessToken');
        let response = await axios.post(`${url}${endpoint}`, body, {
            headers: {
                Authorization: `Bearer ${accessToken}`
            }
        });

        console.log(response.data);

        if(response.status === 401) {
            const refreshResult = await refreshAccessTokenRequest(localStorage.getItem('refreshToken'));
            if(refreshResult) {
                response = await axios.post(`${url}${endpoint}`, {
                    headers: { Authorization: `Bearer ${accessToken}` },
                });
            }
            else {
                return null;
            }
        }
        return response.data;
    } catch (error) {
        console.error('Ошибка POST:', error);
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

export const getUserInfoRequest = async () => {
    return sendPrivateGet('/api/v1/auth/me');
};

export const logoutRequest = async () => {
    const data = {
        refreshToken: localStorage.getItem('refreshToken')
    };
    return sendPrivatePost('/api/v1/auth/logout', data);
};

export const refreshAccessTokenRequest = async () => {
    const data = {
        refreshToken: localStorage.getItem('refreshToken')
    };
    const token = await sendPost('/api/v1/auth/refresh', data);
    console.log('новые токены', token);
    if (token) {
        localStorage.setItem('accessToken', token.accessToken);
        localStorage.setItem('refreshToken', token.refreshToken);
        return true;
    }
    return false;
};