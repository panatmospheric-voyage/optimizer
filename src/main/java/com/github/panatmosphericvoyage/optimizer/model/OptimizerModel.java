package com.github.panatmosphericvoyage.optimizer.model;

public abstract class OptimizerModel {

	public abstract void defineUnit(Identifier identifier, Unit unit);

	public abstract void defineProperty(Identifier identifier, Property property);

	public abstract void defineSummarization(Identifier identifier, Property property);

	public abstract void defineRequirement(Identifier identifier, Requirement requirement);

	public abstract void defineEnumeration(Identifier identifier, Enumeration enumeration);

	public abstract void defineAssembly(Identifier identifier, Assembly assembly);

	public abstract void defineAssemblySubassembly(Assembly assembly, Identifier subassemblyIdentifier, Assembly subassembly);

	public abstract void defineAssemblyProperty(Assembly assembly, Identifier identifier, Property property);

	public abstract void defineAssemblySummarization(Assembly assembly, Identifier identifier, Property property);

	public abstract void defineAssemblyRequirement(Assembly assembly, Identifier identifier, Requirement requirement);

	public abstract void defineAssemblyEnumeration(Assembly assembly, Identifier identifier, Enumeration enumeration);

}
