import styled from "styled-components";

export const ButtonContainer = styled.button`
    display: flex;
    align-items: center;
    justify-content: center;
    width: auto;
    padding: 3rem 5rem;
    background-color:  ${({ backgroundColor }) => backgroundColor};;
    color: ${({ color }) => color};
    font-size: 4rem;
    border-radius: 8rem;
    border: none;
    cursor: pointer;
`