import { configureStore } from '@reduxjs/toolkit'

import { initialAuthState, authReducer } from "./reducers/authReducer"

import { initialAppState, appReducer } from "./reducers/appReducer"

const initialState = {
  authReducer: initialAuthState,
  appReducer: initialAppState,
};

const reducer = {
  authReducer: authReducer,
  appReducer: appReducer,
};

export const store = configureStore({
  reducer: reducer,
  devTools: true,
  preloadedState: initialState,
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;