import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    isLoggedIn: true,
    isLogin: false,
    freeRequestsCount: localStorage.getItem("freeCount") ?? 5,
    answerContent: null,
    profileName: 'кто прочитал тому счастья здоровья'
};

const appSlice = createSlice({
    name: 'counter',
    initialState: initialState,
    reducers: {
        setIsLoggedIn: (state, action) => {
            state.isLoggedIn = action.payload;
        },
        setIsLogin: (state, action) => {
            state.isLogin = action.payload;
        },
        decrementFreeRequestsCount: (state) => {
            state.freeRequestsCount -= 1;
            localStorage.setItem("freeCount", state.freeRequestsCount);
        },
        setAnswerContent: (state, action) => {
            state.answerContent = action.payload;
        },
        setProfileName: (state, action) => {
            state.profileName = action.payload;
        }
    }
});

export const { setIsLoggedIn,
    setIsLogin,
    decrementFreeRequestsCount,
    setAnswerContent,
    setProfileName} = appSlice.actions;
export default appSlice.reducer;