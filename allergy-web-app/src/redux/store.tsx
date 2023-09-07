import { configureStore } from '@reduxjs/toolkit'

import { initialAuthState } from "./reducers/initialAuthState"
import { authReducer } from "./reducers/initialAuthState"

const initialState = {
  authReducer: initialAuthState,
};

const reducer = {
  authReducer: authReducer,
};

export const store = configureStore({
  reducer: reducer,
  devTools: true,
  preloadedState: initialState,
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;