import styled from 'styled-components';

export const InputContainer = styled.div`
    min-width: 90rem;
    display: flex;
    flex-direction: column;
    gap: 3rem;
`;

export const Label = styled.label`
    font-weight: 700;
    font-size: 4rem;
    line-height: 5rem;
    color: var(--color-charcoal)
`;

export const StyledInput = styled.input`
    padding: 4rem;
    border: 2px solid #8f8f8f;
    border-radius: 3rem;
    font-size: 4rem;

    &:focus {
        outline: none;
        border-color: var(--color-grape);
    }
`;