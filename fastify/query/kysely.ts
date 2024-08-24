import { PrismaClient } from '@prisma/client'
import kyselyExtension from 'prisma-extension-kysely';
import type { DB } from '../prisma/generated/types';
import {
  Kysely,
  PostgresAdapter,
  PostgresIntrospector,
  PostgresQueryCompiler,
} from 'kysely';

function createClient() {
  return new PrismaClient().$extends(
    kyselyExtension({
      kysely: (driver) =>
        new Kysely<DB>({
          dialect: {
            createDriver: () => driver,
            createAdapter: () => new PostgresAdapter(),
            createIntrospector: (db) => new PostgresIntrospector(db),
            createQueryCompiler: () => new PostgresQueryCompiler(),
          },
          plugins: [],
        }),
    }),
  );
}

export type Client = ReturnType<typeof createClient>