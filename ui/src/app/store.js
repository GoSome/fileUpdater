import { configureStore } from '@reduxjs/toolkit'
import updaterReducer from '../updaters/updaterSlice'


export const store = configureStore({
  reducer: {
      updater: updaterReducer,
  },
})
