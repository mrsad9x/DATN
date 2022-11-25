function show(event){
	var target = $(event.target.parentElement).next().attr('class');
	target = target.split(" ");
	target = "."+target[0];
	$(target).hasClass("d-none") ? $(target).removeClass("d-none") : $(target).addClass("d-none");
	
}
