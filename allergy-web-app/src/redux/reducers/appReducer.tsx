import {createAction, createReducer} from '@reduxjs/toolkit'
import {RootState} from "../store"

export const initialAppState = {
    userLat: "",
    userLong: "",
};

const setUserLat = createAction<string>("app/setUserLat");
const setUserLong = createAction<string>("app/setUserLong");

export const appReducer = createReducer(initialAppState, (builder) => {
    builder
        .addCase(setUserLat, (state, action) => {
            state.userLat = action.payload;
        })
        .addCase(setUserLong, (state, action) => {
            state.userLong = action.payload;
        });
});

export const getUserLat = (state: RootState) => {
    return state.appReducer.userLat;
}
export const getUserLong = (state: RootState) => {
    return state.appReducer.userLong;
}