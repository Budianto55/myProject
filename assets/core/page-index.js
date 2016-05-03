viewModel.ind = {}; var ind = viewModel.ind;

ind.Save = function(){
	$.ajax({
	  url: "/user/getsave",
	  method: "POST",
	  data: JSON.stringify({"id": "bbb", "nama":"sinta", "kota":"sby"}),
	  dataType: "json",
	  contentType: "json"
	});
}

$(function(){
	ind.Save();
});