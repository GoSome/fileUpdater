import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  value: 0,
  content: "",
  name: "",
  language: "plaintext",
  filePath: "",
}

export const updaterSlice = createSlice({
  name: 'updater',
  initialState,
  reducers: {
    increment: (state) => {
      // Redux Toolkit allows us to write "mutating" logic in reducers. It
      // doesn't actually mutate the state because it uses the Immer library,
      // which detects changes to a "draft state" and produces a brand new
      // immutable state based off those changes
      state.value += 1
    },
    decrement: (state) => {
      state.value -= 1
    },
    setContent: (state,action) => {
        state.content = action.payload ;
    },
    setName: (state,action) => {
        state.name = action.payload
    },
    setFilePath: (state,action) => {
      state.filePath = action.payload
  },
    setLanguage: (state,action) => {
      state.language = action.payload
  },
    incrementByAmount: (state, action) => {
      state.value += action.payload
    },
  },
})

// Action creators are generated for each case reducer function
export const { setLanguage,setContent,setName,setFilePath } = updaterSlice.actions

export default updaterSlice.reducer