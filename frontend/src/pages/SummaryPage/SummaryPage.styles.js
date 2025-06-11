import styled from 'styled-components';

export const Wrapper = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
`;

export const SummaryContainer = styled.div`
    max-width: 280rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    margin-top: 20rem;
`;

export const Title = styled.h1`
    font-size: 64px;
    line-height: 72px;
    font-weight: 700;
    
    max-width: 260rem;
    text-align: center;
    padding-bottom: 27px;
    margin: 0;

    display: inline-block;
    color: transparent;
    background: linear-gradient(to right, #F9D1FF, #6E56CF, #4573F4, #6E56CF, #F9D1FF);
    background-clip: text;
    
    animation: gradient 2s linear infinite;
    background-size: 200% 200%;
`;

export const Subtitle = styled.p`
    font-size: 20px;
    line-height: 32px;
    font-weight: 400;
    color: #FFFFFF;
    opacity: 80%;

    max-width: 260rem;
    text-align: center;
    padding-bottom: 40px;
    margin: 0;
`;