SELECT 

	agoraAgents.id as agentId,
	agoraAgents.userId,
	agoraAgents.userEmail,

	agoraData.link,
	agoraData.name,
	agoraData.details,
	agoraData.category,
	agoraData.middleCategory,
	agoraData.subCategory,
	agoraData.condition,
	agoraData.image,
	agoraData.area,
	agoraData.date

	FROM agoraData
	inner join agoraAgents
	on agoraData.name like '%' || agoraAgents.searchTxt || '%'
	and agoraAgents.category in (agoraData.category,"")
	and (
		agoraAgents.subCategory in (agoraData.middleCategory,"") 
		or 
		agoraAgents.subCategory in (agoraData.subCategory,"") 
	)
	and agoraAgents.condition in (agoraData.condition,"")
	and agoraAgents.area in (agoraData.area,"")
	-- and (
	-- 	agoraData.image != "" 
	-- 	or 
	-- 	agoraAgents.withImage = false
	-- )
	-- where agoraData.processed = false