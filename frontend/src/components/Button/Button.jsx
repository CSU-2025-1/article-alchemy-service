import * as SC from './Button.styles';

export const Button = ({ backgroundColor, color, content, handleClick, variant='primaryButton' }) => {
    return (
        <SC.ButtonContainer backgroundColor={backgroundColor} color={color} onClick={handleClick} variant={variant}>
            {content}
        </SC.ButtonContainer>
    );
};