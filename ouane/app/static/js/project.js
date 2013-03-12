function columnModif() {
	alert("Normalement ca permettra de modif la colonne");
}

function caseAdd() {
	alert("Normalement ca ajoutera une carte a la colonne");
}

$(function() {

$( "#columnCreation" ).button().click(function() {
	var colContainer = $("<div class=\"column\"></div>");
	var columnTitle = $("<button class=\"columnTitle\">Colonne</button>").button().click(function() { columnModif(); });
	var addCard = $("<button class=\"addCard\">Add a card...</button></div></td>").button().click(function() { caseAdd(); });

	colContainer.append(columnTitle);

	var cardsSpace = $("<div class=\"cardsSpace\"></div>").droppable({ hoverClass: "ui-state-active" });
	var cardsList = $("<ul class=\"cardsList\"></ul>").sortable({ connectWith: "ul" }).disableSelection();

	cardsList.appendTo(cardsSpace);
	cardsSpace.appendTo(colContainer);
	colContainer.append(addCard);
	colContainer.appendTo("#columnsList");
})

$(".cardsList").sortable({
	connectWith: "ul"
}).disableSelection();

$( ".cardsSpace" ).droppable({
	hoverClass: "ui-state-active"
});

$( ".columnTitle" ).button().click(function() {
	columnModif();
})

$(".addCard").button().click(function() {
	caseAdd();
})


});
