import styled from 'styled-components';

export const SearchInputContainer = styled.div`
    display: flex;
    justify-content: space-between;
    width: 157rem;
    padding: 3rem 10px 3rem 5rem;
    border-radius: 10rem;
    background-color: #F5F5F5;
    gap: 2rem;
`;

export const SearchInput = styled.input`
    border: none;
    font-size: 4rem;
    width: 100%;
    background-color: #F5F5F5;
    &:focus {
        outline: none;
    }
    color: var('--color-grape');
`;
