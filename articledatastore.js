use ffdatabase;
var bulk = db.articlesStore.initializeUnorderedBulkOp();
bulk.insert( {_id:1, title:"latest science show that potato chips are better for you than sugar.", date:"2016-09-22", body:"some text, potentially containing simple markup about how potato chips are great.", tags:["health", "fitness", "science"] } );
bulk.execute();