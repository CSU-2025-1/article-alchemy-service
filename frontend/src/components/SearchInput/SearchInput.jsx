import * as SC from './SearchInput.styles'
import { Button } from '@/components/Button';

export const SearchInput = ( { backgroundColorButton, colorButton, contentButton } ) => {
    return (
        <SC.SearchInputContainer>
            <SC.SearchInput placeholder={'Ссылка на статью...'} />
            <Button backgroundColor={backgroundColorButton} color={colorButton} content={contentButton} />
        </SC.SearchInputContainer>
    );
};