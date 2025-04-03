import styled from 'styled-components';

export const AuthContainer = styled.div`
    max-width: 400px;
    margin: 0 auto;
    padding: 12rem;
    background-color: white;
    border-radius: 10rem;
`;

export const AuthForm = styled.form`
    display: flex;
    flex-direction: column;
    gap: 8rem;
    color: var(--color-charcoal)
`;

export const AuthFormTitle = styled.span`
    font-size: 12rem;
    line-height: 15rem;
    font-weight: 700;
`;


export const AuthFormSubtitle = styled.span`
    font-size: 4rem;
    line-height: 6rem;
    font-weight: 400;
    color: #667085;
`;

export const ToggleText = styled.p`
    display: flex;
    gap: 2rem;
    font-size: 3rem;
    line-height: 5rem;
    color: var(--color-cadetGray);
`;

export const ToggleButton = styled.button`
    background: none;
    border: none;
    color: var(--color-grape);
    cursor: pointer;
    font-weight: bold;
    
    &:hover {
        text-decoration: underline;
    }
`;