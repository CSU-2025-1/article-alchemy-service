import { Button } from '@/components/Button';
import { SearchInput } from '@/components/SearchInput';

export const SummaryPage = () => {
    return (
        <div>
            <SearchInput backgroundColorButton={'var(--color-grape)'} colorButton={'white'} contentButton={'Кнопка'}> </SearchInput>
            <Button backgroundColor={'var(--color-grape)'} color={'white'} content={'Кнопка'}></Button>
        </div>
    );
};