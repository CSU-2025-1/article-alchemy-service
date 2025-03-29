import styled, {css} from "styled-components";

const ButtonVariants = {
    primaryButton: css`
        border-radius: 5rem;
        padding: 3rem 5rem;
    `,
    authButton: css`
        border-radius: 3rem;
        padding: 4.5rem 5rem;
        font-weight: 700;
    `,
};


export const ButtonContainer = styled.button`
    display: flex;
    align-items: center;
    justify-content: center;
    width: auto;
    background-color:  ${({ backgroundColor }) => backgroundColor || 'var(--color-grape)'};
    color: ${({ color }) => color || 'white'};
    font-size: 4rem;
    border: none;
    cursor: pointer;
    ${({ variant }) => ButtonVariants[variant]};
`