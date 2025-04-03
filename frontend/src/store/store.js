import {configureStore} from "@reduxjs/toolkit";
import appSlice from "@/store/appSlice.js";

export const store = configureStore({
    reducer: appSlice,
});