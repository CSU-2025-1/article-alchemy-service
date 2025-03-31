import styled from "styled-components";

export const HistoryContainer = styled.div`
    display: flex;
    flex-direction: column;
    gap: 10rem;
    align-items: center;
`;

export const HistoryElementContainer = styled.div`
    padding: 3.5rem 5rem 3.5rem 7.5rem;
    display: flex;
    justify-content: space-between;
    width: 156.5rem;
    border-radius: 50px;
    background-color: #FFFFFF;
    
    &:first-child {
        margin-top: 12.25rem;
    }
`;

export const HistoryTitle = styled.h2`
    font-size: 4rem;
`;