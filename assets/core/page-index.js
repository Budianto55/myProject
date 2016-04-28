viewModel.ind = {}; var ind = viewModel.ind;

$(function(){
	$.ajax({
	  url: "/user/getsave",
	  method: "POST",
	  data: JSON.stringify({"id": "www", "nama":"ccc", "kota":"sby"}),
	  dataType: "json",
	  contentType: "json"
	})
});