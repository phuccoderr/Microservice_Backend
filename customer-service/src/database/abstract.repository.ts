import { AbstractDocument } from './abstract.schema';
import { FilterQuery, Model } from 'mongoose';
import { NotFoundException, Logger } from '@nestjs/common';

export abstract class AbstractRepository<TDocument extends AbstractDocument> {

  protected abstract readonly logger: Logger;

  constructor(protected readonly model: Model<TDocument>) {}

  async findOne(
    filterQuery: FilterQuery<TDocument>,
    select: string
  ): Promise<TDocument> {
    const document = await this.model.findOne(filterQuery).select(select).lean<TDocument>(true);

    if (!document) {
      this.logger.warn('Document was not found with filterQuery', filterQuery);
      throw new NotFoundException('Document was not found');
    }

    return document;
  }

  async find(filterQuery: FilterQuery<TDocument>): Promise<TDocument[]> {
    return this.model.find(filterQuery).lean<TDocument[]>(true);
  }
}