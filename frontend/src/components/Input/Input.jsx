import * as SC from './Input.styles';

export const Input = ({ label, ...props }) => {
    return (
        <SC.InputContainer>
            <SC.Label>{label}</SC.Label>
            <SC.StyledInput {...props} />
        </SC.InputContainer>
    );
};