import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    isLoggedIn: false,
    isLogin: false,
    freeRequestsCount: parseInt(localStorage.getItem("freeCount")) ?? 5,
    answerContent: {body: null, status: 'none'},
    profileName: '',
    profileEmail: '',
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
        },
        setProfileEmail: (state, action) => {
            state.profileEmail = action.payload;
        }
    }
});

export const { setIsLoggedIn,
    setIsLogin,
    decrementFreeRequestsCount,
    setAnswerContent,
    setProfileName,
    setProfileEmail} = appSlice.actions;
export default appSlice.reducer;