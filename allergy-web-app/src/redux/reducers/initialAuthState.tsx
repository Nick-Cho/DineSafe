import { createAction, createReducer } from '@reduxjs/toolkit'
import { RootState } from '../store'

export const initialAuthState = {
  userId: "",
  userEmail: ""
}

const setUserId = createAction<string>("auth/setUserId");
const setUserEmail = createAction<string>("auth/setUserEmail");

export const authReducer = createReducer(initialAuthState, (builder) => {
  builder
    .addCase(setUserId, (state, action) => {
      state.userId = action.payload;
    })
    .addCase(setUserEmail, (state, action) => {
      state.userEmail = action.payload;
    });
});

export const getUserId = (state: RootState) => {
  return state.authReducer.userId;
}

export const getUserEmail = (state: RootState) => {
  return state.authReducer.userEmail;
}