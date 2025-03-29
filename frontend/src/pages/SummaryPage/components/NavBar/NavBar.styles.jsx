import styled from 'styled-components';

export const NavContainer = styled.div`
    width: 100%;
`;
export const NavBar = styled.div`
    max-width: 280rem;
    height: 20rem;
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 2px solid var(--color-gray);
`;
export const NavElem = styled.div`
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 3rem;
`;