const TextArea = (props) => {
	return (
		<div className="mb-3">
			<label htmlFor={props.name} className="form-label">
				{props.title}
			</label>
			<textarea rows={props.rows} className="form-control" id={props.name} name={props.name} value={props.value}
				   onChange={props.handleChange} placeholder={props.placeholder}/>
		</div>
	);
};

export default TextArea;