type Action = {
    type: "setId"
    payload: string
}

const reducer = (state: number = 0, action: Action) => {
    return action.payload;
}

export default reducer
