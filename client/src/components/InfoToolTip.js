import oopsImage from '../images/oops.svg';

function InfoToolTip(props) {
    const isOpen = props.isOpen ? 'popup_opened' : '';

    return (
        <div onMouseDown={(e) => {
            if (e.target === e.currentTarget) {
                props.onClose();
            }
        }} className={`popup ${isOpen}`}>
            <div className='popup__container'>
                <button onMouseDown={props.onClose} className='close-button close-button_type_info-tooltip button'></button>
                <img className='popup__image' src={oopsImage} alt='sad face'/>
                <h2 className='popup__title'>Ooops!</h2>
                <p className='popup__subtitle'>Что-то пошло не так! Попробуйте <br></br> ещё раз.</p>
                <button onMouseDown={props.onClose} className='submit-button'>Попробовать ещё</button>
            </div>
        </div>
    );
}

export default InfoToolTip;